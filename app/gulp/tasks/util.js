'use strict';

var gulp = require('gulp');
var gUtil = require('gulp-util');

module.exports = {
  info: function(type, environment, version, minified) {
    gUtil.log('--------------------------------');
    gUtil.log('|  Type: ', gUtil.colors.green(type));
    gUtil.log('|');
    gUtil.log('|  Environment: ', gUtil.colors.yellow(environment));
    gUtil.log('|  Version:     ', gUtil.colors.yellow(version));
    if (minified) {
      gUtil.log('|  Minified');
    } else {
      gUtil.log('|  Unminified');
    }
    gUtil.log('--------------------------------');
  },
  clean: function(dirs, cb) {
    var del = require('del');
    del(dirs).then(function() {
      cb();
    });
  }
};
