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
      MOMO_WSURL: '"ws://192.168.2.5:8080/ws"',
      SW_WSURL: '"wss://plarail2021-py.gotti.dev/"',
    }),
  ],
};
