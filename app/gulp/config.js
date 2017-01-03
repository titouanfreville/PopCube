'use strict';

const gUtil = require('gulp-util');

const urls = {
  development: 'https://development.popcube.com',
  staging: 'https://staging.popcube.com',
  sandbox: 'https://sandbox.popcube.com',
  production: 'https://www.popcube.com'
};

const target = gUtil.env.target || 'development';
const dir = gUtil.env.dir || 'build';
const url = urls[target];
const version = gUtil.env.version || Date.now();

const conf = {
  target: target,
  build: {dir: dir, target: target, url: url, version: version},
  src: {
    js: [
      'app/scripts/**/*.js', '!app/scripts/**/*.spec.js',
      // exclude old files and directories
      '!app/**/*-old/**', '!app/**/*-old.js'
    ],
    testFiles: ['app/scripts/**/*.spec.js'],
    sassBuilds: ['app/scripts/popcube.web.scss'],
    sass: 'app/**/*.scss',
    html: 'app/**/*.html',
    styleGuide: 'app/scripts/dev/style/guide',
    cachedTemplates: [
      'app/scripts/public/pages/welcome/**/*.html'
      // use this to cache every template
      /*
      'app/scripts/!**!/!*.html',
      'app/views/!**!/!*.html',
      '!app/scripts/dev/!**!/!*.html'
      */
    ]
  },
  tmp: '.tmp',
  imgMin: gUtil.env.img != null ? gUtil.env.img : true,
  dev: gUtil.env.dev != null  ? gUtil.env.dev : true,
  watch: gUtil.env.watch != null ? gUtil.env.watch : true,
  weinre: gUtil.env.weinre != null ? gUtil.env.weinre : false,
  bower: gUtil.env.bower != null ? gUtil.env.bower : true,
  debug: gUtil.env.debug != null ? gUtil.env.debug : false,
  test: gUtil.env.test != null ? gUtil.env.test : true,
  debugInBrowser: gUtil.env['debug-in-browser'] != null ? gUtil.env['debug-in-browser'] : false,
  min: gUtil.env.min != null ? gUtil.env.min : false,
  bench: gUtil.env.bench != null ? gUtil.env.bench : false
};

module.exports = conf;
