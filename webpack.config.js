/**
 *     "start": "rm -rf build/ && webpack-dev-server --env.mode development --hot",
 "build": "rm -rf build/ && webpack --env.mode production"
 * @type {module:path}
 */

const
  path = require("path"),
  webpack = require("webpack"),
  webpackMerge = require("webpack-merge"),
  HtmlWebpackPlugin = require("html-webpack-plugin"),
  ResourcesManifestPlugin = require("resources-manifest-webpack-plugin");

function modeConfig(env) {
  return require(`${path.resolve(__dirname)}/webpack.${env}`)(env);
}

module.exports = ({ mode } = { mode: "production" }) => {
  return webpackMerge({
    mode,
    entry: "./src/index.js",
    optimization: {
      splitChunks: {
        chunks: "all"
      }
    },
    module: {
      rules: [
        {
          test: /\.jsx?$/,
          use: ["babel-loader"],
          exclude: /node_modules/
        },
        {
          test: /\.(png|jpe?g|gif|svg|webp)$/,
          use: [
            {
              loader: "file-loader",
              options: {
                name: "static/media/[name].[contenthash].[ext]"
              }
            }
          ]
        }
      ]
    },
    resolve: {
      extensions: ["*", ".js", ".jsx"],
      alias: {
        src: path.resolve(__dirname, "src/")
      }
    },
    devServer: {
      historyApiFallback: true,
      contentBase: path.join(__dirname, "public")
    },
    plugins: [
      new HtmlWebpackPlugin({
        template: "public/index.html",
        inject: "body",
        minify: {
          html5: true,
          removeComments: true,
          collapseWhitespace: true
        },
        templateParameters: {
          PUBLIC_URL: ""
        }
      }),
      new ResourcesManifestPlugin({
        TO_CACHE: /.+\.(js|css|png|jpe?g|gif|svg|webp)$/
      }, "public/service-worker.js"),
      new webpack.ProgressPlugin()
    ]
  }, modeConfig(mode));
};
