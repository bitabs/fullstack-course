const
  path                      = require('path'),
  MiniCssExtractPlugin      = require('mini-css-extract-plugin'),
  OptimizeCSSAssetsPlugin   = require('optimize-css-assets-webpack-plugin'),
  CopyWebpackPlugin         = require('copy-webpack-plugin'),
  { BundleAnalyzerPlugin }  = require("webpack-bundle-analyzer"),
  UglifyJsPlugin = require('uglifyjs-webpack-plugin');


module.exports = () => ({
  devtool: 'none',
  output: {
    path: path.resolve(__dirname, './build/'),
    publicPath: '/',
    filename: 'static/js/[name].[contenthash].js'
  },
  module: {
    rules: [{
      test: /\.css$/,
      use: [MiniCssExtractPlugin.loader, 'css-loader']
    }]
  },
  optimization: {
    minimizer: [new UglifyJsPlugin()],
  },
  plugins: [
    new MiniCssExtractPlugin({
      filename: 'static/css/[name].[contenthash].css'
    }),
    new CopyWebpackPlugin([{
      from: 'public/',
      to: '.',
      ignore: ['service-worker.js']
    }]),
    // new BundleAnalyzerPlugin(),
    new OptimizeCSSAssetsPlugin()
  ]
})
