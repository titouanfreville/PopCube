/**
 * Copy tasks
 */

'use strict';

const gulp = require('gulp');
const conf = require('./config');

gulp.task('modernizr', function (done) {
  const modernizr = require('modernizr');
  const fs = require('fs');
  const config = require('../app/libs/modernizr-config.json');
  modernizr.build(config, function(code) {
    fs.writeFile(__dirname + '/../app/libs/modernizr.js', code, done);
  });
});

gulp.task('copy:tmp', function() {
  gulp.src([
    'app/libs/**/*',
    'app/angular/**/*',
    'app/styles/css/*.css'
  ], {base: 'app'})
    .pipe(gulp.dest(conf.tmp));
  return gulp.src([
    conf.src.bowerComponents + '/**/*'
  ])
    .pipe(gulp.dest(conf.tmp + '/node_modules'));
});

gulp.task('copy:assets', ['copy:popcube_font'], function() {
  return gulp.src([
    conf.src.bowerComponents + '/font-awesome/fonts/**'
  ])
    .pipe(gulp.dest(conf.build.dir + '/styles/fonts'));
});

gulp.task('copy:popcube_font', function() {
  return gulp.src([
    conf.src.bowerComponents + '/popcube_font/**'
  ])
    .pipe(gulp.dest(conf.build.dir + '/styles/fonts/popcube_font'));
});

gulp.task('copy', ['copy:build', 'copy:popcube_font'], function() {
  if (conf.mobile) {
    const sources = ['native/cobalt.conf','native/start.html'];
    const debug = conf.debugInBrowser;
    if (debug) {
      sources.push('native/cobalt.debug.mobile.js');
    }
    gulp.src(sources).pipe(gulp.dest(conf.build.dir));
    const imgSources = ['app/scripts/public/pages/sign-up/images/**/*.jpg'];
    gulp.src(imgSources, {base: 'app'})
      .pipe(gulp.dest(conf.build.dir));
  }
  gulp.src('app/libs/modernizr.js', {base: 'app/libs'}).pipe(gulp.dest(conf.build.dir + '/scripts'));
  let src = [
    'app/**/*.mp4',
    'app/**/*.css',
    'app/**/*.png',
    'app/**/*.jpg',
    'app/images/**'
  ];
  if (conf.mobile) {
    src = src.concat(['!app/scripts/public/**/*.png', '!app/scripts/public/**/*.jpg', '!app/**/*.mp4']);
  }
  return gulp.src(src, {base: 'app'})
    .pipe(gulp.dest(conf.build.dir));
});

gulp.task('copy:build', function() {
  return gulp.src([
    'app/robots.txt', 'app/favicon.ico', 'app/*.png',
    'app/static/**', 'app/sitemap.xml', 'app/languages/angular-i18n/**',
    'app/apple-app-site-association'
  ], {base: 'app'})
    .pipe(gulp.dest(conf.build.dir));
});
