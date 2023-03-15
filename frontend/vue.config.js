/**
 * @type {import('@vue/cli-service').ProjectOptions}
 */
module.exports = {
  // Remove moment.js from chart.js
  configureWebpack: config => {
    return {
      externals: {
        moment: 'moment'
      },
      resolve: {
        fallback: {
          // assert: require.resolve("assert/"),
          // async_hooks: false,
          // fs: false,
          // path: require.resolve("path-browserify"),
          // http: require.resolve('stream-http'),
          // https: require.resolve('https-browserify'),
          // url: require.resolve('url/'),
          // os: require.resolve("os-browserify/browser"),
          // stream: require.resolve("stream-browserify"),
          // zlib: require.resolve("browserify-zlib"),
          // async_hooks: false,
          // dns: false,
          // "crypto": require.resolve("crypto-browserify")
        }
      }
    }
  }
}
