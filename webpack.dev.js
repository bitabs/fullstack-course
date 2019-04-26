const
  path = require('path'),
  merge = require('webpack-merge'),
  common = require('./webpack.common.js');

module.exports = merge(common, {
  mode: 'development',
  devtool: 'inline-source-map',
  devServer: {
    contentBase: "public/",
    port: 3000,
    hot: true
  }
});