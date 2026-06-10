import vueParser from 'vue-eslint-parser'
import tseslint from 'typescript-eslint'

export default [
  {
    ignores: ['node_modules/**', 'dist/**', 'build/**', 'coverage/**', 'playwright-report/**', 'test-results/**']
  },
  {
    files: ['**/*.ts'],
    languageOptions: {
      parser: tseslint.parser,
      parserOptions: { ecmaVersion: 'latest', sourceType: 'module' }
    }
  },
  {
    files: ['**/*.vue'],
    languageOptions: {
      parser: vueParser,
      parserOptions: {
        parser: tseslint.parser,
        ecmaVersion: 'latest',
        sourceType: 'module'
      }
    }
  }
]
