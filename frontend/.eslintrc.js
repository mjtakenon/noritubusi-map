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
		"no-undef": "off",
		"no-unused-vars": "off"
  },
  parserOptions: {
    parser: "babel-eslint"
  }
};
