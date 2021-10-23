const path = require("path");
const webpack = require("webpack");

module.exports = {
  entry: {
    skyway_receiver: "./skyway_receiver.js",
    momo_sender: "./momo_sender.js",
  },
  output: {
    filename: "[name]_.js",
    path: path.resolve(__dirname),
  },
  plugins: [
    new webpack.DefinePlugin({
      MOMO_WSURL: '"ws://127.0.0.1:8080/ws"',
      SW_WSURL: '"ws://127.0.0.1:8081"',
    }),
  ],
};
