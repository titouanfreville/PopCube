/**
 * Compile tasks
 */

'use strict';

const gulp = require('gulp');
const gUtil = require('gulp-util');
const conf = require('./config');
const webpack = require('webpack');

let webpackConf;
if (conf.mobile) {
  if (conf.min) {
    webpackConf = require('../webpack.config.mobile.build.js');
  }else {
    webpackConf = require('../webpack.config.mobile.js');
  }
} else {
  if (conf.min) {
    webpackConf = require('../webpack.config.build.js');
  } else {
    webpackConf = require('../webpack.config.js');
  }
}

gulp.task("webpack", function(callback) {
  // run webpack
  let called = false;
  webpack(webpackConf, function(err, stats) {
    if(err) {
      throw new gUtil.PluginError('webpack', err);
    }
    gUtil.log('[webpack]', stats.toString('minimal'));
    if (stats.hasErrors() && conf.min) {
      gUtil.log(gUtil.colors.red('Error during webpack build, exiting.'));
      process.exit(1);
    }
    if (!called) {
      callback();
      called = true;
    }
  });
});
