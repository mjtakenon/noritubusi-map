module.exports = {
  // 開発時に動的コンパイルを有効にするためのオプション
  runtimeCompiler: true,
  // 開発環境とデプロイ環境でのURL共通化
  devServer: {
    proxy: {
      // api/*  --> localhost:1323/*
      '.*': {
        target: 'http://localhost:1323',
        pathRewrite: {
          '^/api': ''
        }
      }
    }
  }
};
