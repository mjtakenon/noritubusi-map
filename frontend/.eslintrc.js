module.exports = {
  root: true,
  env: {
    node: true,
    browser: true
  },
  extends: ["plugin:vue/essential", "@vue/prettier"],
  rules: {
    "no-console": process.env.NODE_ENV === "production" ? "error" : "off",
    "no-debugger": process.env.NODE_ENV === "production" ? "error" : "off",

    // 文末セミコロンがない箇所に Warning を出す
    "semi": ["warn", "always"],

    // undefined エラーを非検知に
    "no-undef": "off",

    // 未使用の変数宣言に関するエラーを非検知に
    "no-unused-vars": "off"
  },
  parserOptions: {
    parser: "babel-eslint"
  }
};
