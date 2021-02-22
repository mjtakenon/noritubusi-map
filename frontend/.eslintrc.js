module.exports = {
  root: true,

  env: {
    node: true,
    browser: true,
  },

  extends: ["plugin:vue/essential", "@vue/prettier"],

  rules: {
    "prettier/prettier": [
      "warn",
      {
        // 文末セミコロンに関するエラー
        semi: false,
        trailingComma: "es5",
      },
    ],
    "no-console": process.env.NODE_ENV === "production" ? "error" : "off",
    "no-debugger": process.env.NODE_ENV === "production" ? "error" : "off",

    // undefined エラー
    "no-undef": "off",

    // 未使用の変数宣言に関するエラー
    "no-unused-vars": "off",

    // Vue: import したが未使用のコンポーネントに対するエラー
    "vue/no-unused-components": "off",
  },

  parserOptions: {
    parser: "@typescript-eslint/parser",
  },

  extends: ["plugin:vue/essential", "@vue/prettier", "@vue/typescript"],
}
