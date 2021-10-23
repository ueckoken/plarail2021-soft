const path = require("path");

module.exports = {
  entry: {
    skyway_receiver: "./skyway_receiver.js",
    momo_sender: "./momo_sender.js",
  },
  output: {
    filename: "[name]_.js",
    path: path.resolve(__dirname),
  },
};
