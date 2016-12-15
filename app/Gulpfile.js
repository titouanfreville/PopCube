/* jshint node: true */
'use strict';

// Require build dependencies
const gulp = require('gulp');
const gUtil = require('gulp-util');
const conf = require('./gulp/config');
const util = require('./gulp/tasks/util');

gulp.task('info', function() {
  if (conf.mobile) {
    util.info('mobile', conf.build.target, conf.build.version);
  } else {
    util.info('web', conf.build.target, conf.build.version);
  }
});

/**
 * Update submodules
 */
gulp.task('git:update-submodules', function(done) {
  if (conf.mobile) {
    const git = require('gulp-git');
    git.updateSubmodule({args: '--init'}, done);
  } else {
    done();
  }
});

/**
 * Clean and dependencies tasks
 */
gulp.task('clean', function(cb) {
  if (conf.mobile) {
    util.clean([
      conf.mobileBuild.dir,
      conf.mobileBuild.tmp,
      conf.src.bowerComponents
    ], cb);
  } else {
    util.clean([
      conf.tmp,
      conf.build.dir,
      conf.src.bowerComponents,
      '.tests'
    ], cb);
  }
});

gulp.task('bower', function(done) {
  const bower = require('bower');
  if (!conf.bower) {
    gUtil.log(['bower is running an', gUtil.colors.yellow('offline'), 'install'].join(' '));
  }
  bower.commands.install([], {}, {offline: !conf.bower})
  .on('error', function(error) {
    gUtil.log(['bower', gUtil.colors.red(error)].join(' '));
    process.exit(1);
  })
  .on('end', function() {
    done();
  });
});

require('./gulp/gulp_replace');
require('./gulp/gulp_webpack');
require('./gulp/gulp_minify');
require('./gulp/gulp_dev');
require('./gulp/gulp_copy');
require('./gulp/gulp_linters');
require('./gulp/gulp_tests');

/**
 * CLI tasks
 */

const rseq = require('run-sequence');

// Development task
gulp.task('default', function(cb) {
  rseq(
  'info', 'clean',
  ['bower', 'git:update-submodules'],
  ['replace', 'copy:bower', 'guide'],
  'webpack',
  'inject',
  'copy',
  // 'replace:tests',
  // 'test:unit:dev:once',
  // 'watch',
  cb);
});


gulp.task('test', function(cb) {
  conf.imgMin = false;
  rseq(
  'build',
  'test:unit',
  cb);
});
