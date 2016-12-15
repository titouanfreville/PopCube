/**
 * Tests tasks
 */

'use strict';

const gulp = require('gulp');
const gUtil = require('gulp-util');
const conf = require('./config');

function runKarma(configFilePath, cb, dev) {
  const Server = require('karma').Server;
  const path = require('path');
  const configPath = path.resolve(configFilePath);

  const server = new Server({configFile: configPath}, function(exitCode) {
    gUtil.log('Karma has exited with ' + gUtil.colors[exitCode ? 'red' : 'green'](exitCode));
    if (exitCode !== 0 && !dev) {
      process.exit(exitCode);
    }
    return cb();
  });
  server.start();
}
gulp.task('test:compile', function() {
  const babel = require('gulp-babel');
  return gulp.src(conf.src.testFiles)
    .pipe(babel({presets: ['es2015']}))
    .pipe(gulp.dest('.tests'));
});
gulp.task('test:unit', ['test:compile'], function(cb) {
  if (conf.test) {
    return runKarma('test/karma-build.conf.js', cb);
  } else {
    return cb();
  }
});
gulp.task('test:unit:dev', ['test:compile'], function(cb) {
  if (conf.test) {
    return runKarma('test/karma-dev-ci.conf.js', cb, true);
  } else {
    return cb();
  }
});
gulp.task('test:unit:dev:once', ['test:compile'], function(cb) {
  if (conf.test) {
    return runKarma('test/karma-dev-once.conf.js', cb, false);
  } else {
    return cb();
  }
});

gulp.task('test:e2e', function() {
  const protractor = require('gulp-protractor').protractor;

  gulp.src(['tests/e2e/**/*.js'])
  .pipe(protractor({
    configFile: 'test/e2e/protractor.config.js',
    args: ['--baseUrl', conf.build.url],
    seleniumServerJar: './node_modules/protractor/selenium/selenium-server-standalone-2.51.0.jar'
  }))
  .on('error', function(e) {
    throw e;
  });
});

gulp.task('fabien', ['test:e2e'], function() {
});
