/**
 * Minification tasks
 */

'use strict';

const gulp = require('gulp');
const inject = require('gulp-inject');
const replace = require('gulp-replace-task');
const conf = require('./config');

const angularVersion = require('../package.json').dependencies.angular;
const devLibs = '\t<script src="https://ajax.googleapis.com/ajax/libs/angularjs/' + angularVersion + '/angular.js"></script>\n' +
  '\t<script src="https://ajax.googleapis.com/ajax/libs/angularjs/' + angularVersion + '/angular-resource.js"></script>\n' +
  '\t<script src="https://ajax.googleapis.com/ajax/libs/angularjs/' + angularVersion + '/angular-cookies.js"></script>\n' +
  '\t<script src="https://ajax.googleapis.com/ajax/libs/angularjs/' + angularVersion + '/angular-animate.js"></script>\n' +
  '\t<script src="https://ajax.googleapis.com/ajax/libs/angularjs/' + angularVersion + '/angular-sanitize.js"></script>\n' +
  '\t<script src="https://cdnjs.cloudflare.com/ajax/libs/socket.io/1.4.5/socket.io.js"></script>';

const prodLibs = '\t<script src="https://ajax.googleapis.com/ajax/libs/angularjs/' + angularVersion + '/angular.min.js"></script>\n' +
  '\t<script src="https://ajax.googleapis.com/ajax/libs/angularjs/' + angularVersion + '/angular-resource.min.js"></script>\n' +
  '\t<script src="https://ajax.googleapis.com/ajax/libs/angularjs/' + angularVersion + '/angular-cookies.min.js"></script>\n' +
  '\t<script src="https://ajax.googleapis.com/ajax/libs/angularjs/' + angularVersion + '/angular-animate.min.js"></script>\n' +
  '\t<script src="https://ajax.googleapis.com/ajax/libs/angularjs/' + angularVersion + '/angular-sanitize.min.js"></script>\n' +
  '\t<script src="https://cdnjs.cloudflare.com/ajax/libs/socket.io/1.4.5/socket.io.min.js"></script>';

gulp.task('inject', function() {
  const cobaltFile = conf.debugInBrowser ? '<script src="cobalt.debug.mobile.js"></script>' : '<script src="../platform/cobalt.min.js"></script>';

  let indexPath;
  let dest;
  if (conf.mobile) {
    indexPath = ['native/index.html', 'native/sidemenu.html'];
    dest = 'native/build';
  } else {
    dest = conf.build.dir;
    if (conf.min) {
      indexPath = '.tmp/index.html';
    } else {
      indexPath = 'app/index.html';
    }
  }
  return gulp.src(indexPath)
    .pipe(replace({
      patterns: [
        {
          match: 'thirdPartyLibs',
          replacement: conf.min ? prodLibs : devLibs
        },
        {
          match: 'CobaltInclusionHere',
          replacement: cobaltFile
        }
      ]
    }))
    .pipe(gulp.dest(dest));
});

gulp.task('imagemin', function() {
  let stream;
  if (!conf.imgMin) {
    stream = gulp.src(['app/images/**/*', 'app/scripts/**/*.png', 'app/scripts/**/*.jpg', 'app/scripts/**/*.gif', 'app/scripts/**/*.jpeg', 'app/scripts/**/*.svg'], {base: 'app'})
      .pipe(gulp.dest(conf.build.dir));
  } else {
    const imagemin = require('gulp-imagemin');
    stream = gulp.src(['app/images/**/*', 'app/scripts/**/*.png', 'app/scripts/**/*.jpg', 'app/scripts/**/*.gif', 'app/scripts/**/*.jpeg', 'app/scripts/**/*.svg'], {base: 'app'})
      .pipe(imagemin({progressive: true, optimizationLevel: 2}))
      .pipe(gulp.dest(conf.build.dir));
  }
  return stream;
});

gulp.task('jsonmin', function() {
  const jsonmin = require('gulp-jsonminify');
  return gulp.src(['app/languages/*.json'], {base: 'app'})
    .pipe(jsonmin())
    .pipe(gulp.dest(conf.build.dir));
});
