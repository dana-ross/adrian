var nodeExternals = require('webpack-node-externals');

module.exports = {
    target: 'node',
    externals: [nodeExternals()], // in order to ignore all modules in node_modules folder    
    context: __dirname + "/app/src",
    entry: "./server.js",
    output: {
        path: __dirname + "/app/dist",
        filename: "bundle.js"
    },
    module: {
        rules: [
            {enforce: 'post', test: /fontkit[\/\\]index.js$/, loader: "transform-loader?brfs"},
            {enforce: 'post', test: /unicode-properties[\/\\]index.js$/, loader: "transform-loader?brfs"},
            {enforce: 'post', test: /linebreak[\/\\]src[\/\\]linebreaker.js/, loader: "transform-loader?brfs"}
        ]
    }
};
