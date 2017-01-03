/**
 * Replace tasks
 */

'use strict';

const gulp = require('gulp');
const conf = require('./config');

/**
 * Create the app configuration depending on build target
 */

gulp.task('replace', function() {
  // util.configFile(conf.build.dir + '/scripts/config', conf.tmp + '/scripts/config', {dev: true, debugInBrowser: false});
});

gulp.task('replace:build', function() {
  // util.configFile(conf.build.dir + '/scripts/config', conf.tmp + '/scripts/config', {dev: false, debugInBrowser: false});
});

function replaceForTests(dir) {
  const fs = require('fs');
  const replace = require('gulp-replace-task');
  return gulp.src('test/confs/*.js')
 /*   .pipe(replace({
      patterns: [
        {
          match: /\/\/ @@inject/,
          replacement: function() {
            const index = fs.readFileSync(conf.build.dir + '/index.html').toString();
            const total = '';
            const match;
            const reg = /<script src="(\/(?:scripts|bower_components|angular|libs).*)"><\/script>/g;
            match = reg.exec(index);
            while (match !== null) {
              total += '\n      \'' + dir + match[1] + '\',';
              match = reg.exec(index);
            }
            return total;
          }
        }
      ]
    }))*/
    .pipe(gulp.dest('test'));
}

gulp.task('replace:tests', function() {
  return replaceForTests(conf.build.dir);
});

gulp.task('replace:tests:build', function() {
  return replaceForTests(conf.tmp);
});
