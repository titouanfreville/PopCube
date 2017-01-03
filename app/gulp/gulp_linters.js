/**
 * Lint tasks
 */

'use strict';

const gulp = require('gulp');
const gUtil = require('gulp-util');

const analyseLangObject = function(lang1Key, lang1Value, lang2Key, lang2Value, results, parent) {
  if (typeof lang1Value === 'object') {
    if (typeof lang2Value !== 'object' || lang2Value === null) {
      results.push(parent + '.' + lang1Key);
    } else {
      parent = parent ? parent + '.' : '';
      for (const keyObj in lang1Value) {
        if (lang1Value.hasOwnProperty(keyObj)) {
          analyseLangObject(keyObj, lang1Value[keyObj], keyObj, lang2Value[keyObj], results, parent + lang2Key);
        }
      }
    }
  } else if (typeof lang2Value === 'undefined' || lang2Value === null || lang2Value === '') {
    results.push(parent + '.' + lang1Key);
  }
  return results;
};

// TODO: Check that keys are unique
gulp.task('languages', function(cb) {
  const en = require('../app/languages/en.json');
  const fr = require('../app/languages/fr.json');

  // Keys not in both i18n files
  const errors = []
    .concat(analyseLangObject('FR', fr, 'EN', en, [], ''))
    .concat(analyseLangObject('EN', en, 'FR', fr, [], ''));
  errors.forEach(function(key) {
    gUtil.log(gUtil.colors.yellow(key + ' translation does not exist'));
  });

  if (errors.length) {
    process.exit(1);
  }

  cb();
});
