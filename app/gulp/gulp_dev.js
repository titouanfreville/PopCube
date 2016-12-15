/**
 * Compile tasks
 */

'use strict';

const gulp = require('gulp');
const marked = require('gulp-markdown');
const replace = require('gulp-replace-task');
const fs = require('fs');
const path = require('path');
const conf = require('./config');
const marked2 = require('marked');
const renderer = new marked2.Renderer();
const rename = require('gulp-rename');
const cheerio = require('cheerio');
const rseq = require('run-sequence');

const guideFiles = conf.src.styleGuide;

let codeIndex = 0;

function escape(text) {
  let string;
  if (text) {
    string = text.toLowerCase().replace(/[^\w]+/g, '-');
  } else {
    string = '';
  }
  return string;
}

renderer.heading = function(text, level) {
  const escapedText = escape(text);

  return '<h' + level + ' title="' + text + '" class="styleguide-heading"><a id="' + escapedText + '" ui-sref="{\'#\':\'' + escapedText +
    '\'}" ui-sref-opts="{reload:true}" href="' + escapedText + '"><span class="header-link">' +
    text + '</span></a>' + '</h' + level + '>';
};

function generateClickFunction(index, value) {
  return `ctrl.tab${index}=ctrl.tab${index}===${value} ? 0 : ${value}`;
}

function codeHeaders(langs, codeIndex) {
  let header = `<ul  class="langs">
      <li ng-click="${generateClickFunction(codeIndex, 1)}" ng-class="{selected: ctrl.tab${codeIndex}==1}">html</li>`;
  if (langs.indexOf('scss') >= 0) {
    header += `<li ng-click="${generateClickFunction(codeIndex, 2)}" ng-class="{selected: ctrl.tab${codeIndex}==2}">scss</li>`;
  }
  if (langs.indexOf('js') >= 0) {
    header += `<li ng-click="${generateClickFunction(codeIndex, 3)}" ng-class="{selected: ctrl.tab${codeIndex}==3}">js</li>`;
  }
  header += '</ul>';
  return header;
}

function codeBlock(lang, code, id, codeIndex) {
  const highlight = require('highlight.js');
  return '<div class="code-block" ng-show="ctrl.tab' + codeIndex + '==' + id + '">' +
    '<pre><code class="lang-' + lang + '">' + highlight.highlightAuto(code, [lang]).value + '</code></pre>' +
  '</div>';
}

gulp.task('markdown', function() {
  return gulp.src('app/famicity.md')
    // include .md files in the main .md file
    .pipe(replace({
      patterns: [
        {
          match: /(?:@include ([^\.]*.md);)/g,
          replacement: function(pattern, file) {
            return fs.readFileSync('app' + path.sep + file, 'utf-8').toString();
          }
        }
      ]
    }))
    // replace html and javascript and scss pieces of code to add the result of this code
    .pipe(replace({
      patterns: [
        {
          match: /(?:%% result %%(?:\n|\r\n)```(html)?(?:\n|\r\n)([^`]+)```\n```(\w*)(?:\n|\r\n)([^`]+)```\n```(\w*)(?:\n|\r\n)([^`]+)```)/g,
          replacement: function(match, p1, p2, p3, p4, p5, p6) {
            codeIndex++;
            return '<div class="guide-result">\n' + p2 + '</div>' +
            '<div class="code-blocks">' +
              codeHeaders(['html', p3, p5], codeIndex) +
              codeBlock('html', p2, 1, codeIndex) +
              codeBlock(p3, p4, 2, codeIndex) +
              codeBlock(p5, p6, 3, codeIndex) +
            '</div>';
          }
        }
      ]
    }))
    // replace html and javascript or scss pieces of code to add the result of this code
    .pipe(replace({
      patterns: [
        {
          match: /(?:%% result %%(?:\n|\r\n)```(html)?(?:\n|\r\n)([^`]+)```\n```(\w*)(?:\n|\r\n)([^`]+)```)/g,
          replacement: function(match, p1, p2, p3, p4) {
            codeIndex++;
            return '<div class="guide-result">\n' + p2 + '</div>' +
            '<div class="code-blocks">' +
              codeHeaders(['html', p3], codeIndex) +
              codeBlock('html', p2, 1, codeIndex) +
              codeBlock(p3, p4, 2, codeIndex) +
            '</div>';
          }
        }
      ]
    }))
    // replace html pieces of code to add the result of this code
    .pipe(replace({
      patterns: [
        {
          match: /(?:%% result %%(?:\n|\r\n)```(html)?(?:\n|\r\n)([^`]+)```)/g,
          replacement: function(match, p1, p2) {
            codeIndex++;
            return '<div class="guide-result">\n' + p2 + '</div>' +
            '<div class="code-blocks">' +
              codeHeaders(['html'], codeIndex) +
              codeBlock('html', p2, 1, codeIndex) +
            '</div>';
          }
        }
      ]
    }))
    .pipe(gulp.dest(guideFiles + '/.tmp'))
    // markdownify
    .pipe(marked({
      highlight: function(code, lang) {
        return require('highlight.js').highlightAuto(code, [lang]).value;
      },
      renderer: renderer,
      breaks: false
    }))
    // include the result of the markdowned content in the main html file
    .pipe(replace({
      patterns: [
        {
          match: /(?:@include ([^\.]*.html);)/g,
          replacement: function(pattern, file) {
            return fs.readFileSync('app' + path.sep + file, 'utf-8').toString();
          }
        }
      ]
    }))
    .pipe(gulp.dest(guideFiles));
});

gulp.task('generate', function() {
  return gulp.src(guideFiles + '/style-guide_template.html')
    .pipe(replace({
      patterns: [
        {
          match: /(?:@include ([^\.]*.html);)/g,
          replacement: function(pattern, file) {
            return fs.readFileSync(guideFiles + path.sep + file, 'utf-8').toString();
          }
        }
      ]
    }))
    .pipe(rename({basename: 'style-guide'}))
    .pipe(gulp.dest(guideFiles));
});

function generateLink(title) {
  const escaped = escape(title);
  return '<a ui-sref="{\'#\': \'' + escaped + '\'}" ui-sref-opts="{reload: true}">' + title + '</a>';
}

gulp.task('menu', function() {
  return gulp.src(guideFiles + '/style-guide.html')
    .pipe(replace({
      patterns: [
        {
          match: /(?:@menu;)/,
          replacement: function() {
            const guide = fs.readFileSync(guideFiles + path.sep + 'style-guide.html', 'utf-8').toString();
            const $ = cheerio.load(guide);
            let menu = '';
            $('h1.styleguide-heading, h2.styleguide-heading').each(function(i, el) {
              if (el.name.toUpperCase() === 'H1') {
                if (i === 0) {
                  menu += '<li>' + generateLink($(el).attr('title')) + '<ul class="nav">';
                } else {
                  menu += '</ul></li><li>' + generateLink($(el).attr('title')) + '<ul class="nav">';
                }
              } else if (i === $('h1.styleguide-heading, h2.styleguide-heading').length - 1) {
                menu += '<li>' + generateLink($(el).attr('title')) + '</li></ul>';
              } else {
                menu += '<li>' + generateLink($(el).attr('title')) + '</li>';
              }
            });
            return menu;
          }
        }
      ]
    }))
    .pipe(gulp.dest(guideFiles));
});

gulp.task('guide:css', function() {
  const plumber = require('gulp-plumber');
  const sourceMaps = require('gulp-sourcemaps');
  const sass = require('gulp-sass');
  const autoprefixer = require('gulp-autoprefixer');
  return gulp.src('app/scripts/dev/style/guide/style-guide.scss')
    .pipe(plumber())
    .pipe(sourceMaps.init())
    .pipe(sass())
    .pipe(autoprefixer({
      browsers: ['last 2 versions', '> 1%']
    }))
    .pipe(sourceMaps.write())
    .pipe(gulp.dest(conf.build.dir + '/scripts/dev/style/guide'));
});

gulp.task('guide', function(cb) {
  rseq(
    ['guide:css', 'markdown'], 'generate', 'menu',
    cb);
});
