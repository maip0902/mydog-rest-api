const MODE = "development";
const path = require('path');
const MiniCssExtractPlugin = require("mini-css-extract-plugin");
const enabledSourceMap = MODE === "development";
 
module.exports = {
  entry: './index.js',  
  output: { 
    filename: 'main.js',
    path: path.resolve(__dirname, 'dist'),
  },
  module: {
    rules: [
      {
        test: /\.scss$/,
        use: [
          {
            loader: MiniCssExtractPlugin.loader,
          },
          {
            loader: "css-loader",
            options: {
              url: false,
              sourceMap: true,
            },
          },
          {
            loader: "sass-loader",
            options: {
              sourceMap: enabledSourceMap
            },
          },
        ],
      },
    ],
  },
  plugins: [
    new MiniCssExtractPlugin({
      filename: "app.css",
    }),
  ],
  devtool: "source-map",
  watchOptions: {
    ignored: /node_modules/  
  },
};