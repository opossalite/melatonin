import adapter from '@sveltejs/adapter-static';
import sveltePreprocess from 'svelte-preprocess';

const config = {
  kit: {
    adapter: adapter({
      fallback: 'index.html',
      pages: 'build',
      assets: 'build',
    }),
    paths: {
      base: '',
    },
  },
  preprocess: sveltePreprocess(),
};

export default config;

