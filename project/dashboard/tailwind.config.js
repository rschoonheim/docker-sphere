/** @type {import('tailwindcss').Config} */
export default {
  content: [
    './index.html',
    './src/**/*.{vue,js,ts,jsx,tsx}',
    './src/views/**/*.{vue,js,ts,jsx,tsx}'
  ],
  safelist: [
    // pattern
    {
      pattern: /h-[(32)]/,
    },
  ],
  theme: {
    extend: {
      fontFamily: {
        'sans': ['IBM Plex Sans', 'sans-serif'], // Primary UI font
        'serif': ['IBM Plex Serif', 'serif'], // Secondary UI font
        'mono': ['IBM Plex Mono', 'monospace'] // For code, numbers, etc.
      },
      colors: {
        'background': {
          DEFAULT: '#161616'
        },
        'layer-01': '#262626',
        'layer-02': '#393939',
        'layer-03': '#525252',
        'layer-hover-01': '#333333',
        'layer-hover-02': '#474747',
        'layer-hover-03': '#636363',

        // Field
        'field-01': '#262626',
        'field-02': '#393939',
        'field-03': '#525252',
        'field-hover-01': '#333333',
        'field-hover-02': '#474747',
        'field-hover-03': '#636363',

        // Border tile
        'border-tile-1': '#c6c6c6',

        // Expanding Text Colors
        'text-primary': '#f4f4f4', // Primary text color
        'text-secondary': '#c6c6c6', // Secondary text color
        'text-tertiary': '#6f6f6f', // Tertiary text color
        'text-quaternary': '#525252', // Quaternary text color
        'text-disabled': '#3d3d3d', // Disabled text color
        'text-inverse': '#161616', // Inverse text color for use on light backgrounds
        'text-on-color': '#ffffff', // Text color for texts on colored backgrounds
        'text-on-color-disabled': '#8d8d8d', // Disabled text color on colored backgrounds
        'text-link': '#0f62fe', // Link text color
        'text-success': '#42be65', // Success state text color
        'text-warning': '#f1c21b', // Warning state text color
        'text-error': '#fa4d56', // Error state text color
        'text-info': '#0f62fe', // Informational text color

        // Retaining previous configurations for completeness
        'interactive-01': '#0f62fe',
        'interactive-02': '#393939',
        'interactive-03': '#0f62fe',
        'interactive-04': '#0f62fe',
        'border-subtle-00': '#8d8d8d',
        'border-subtle-01': '#6f6f6f',
        'border-strong': '#393939',
        'support-success': '#42be65',
        'support-error': '#fa4d56',
        'support-warning': '#f1c21b',
        'support-info': '#0f62fe',
        'link-primary': '#0f62fe',
        'link-secondary': '#78a9ff',
        'link-visited': '#be95ff',
        'icon-primary': '#f4f4f4',
        'icon-secondary': '#c6c6c6',
        'icon-tertiary': '#6f6f6f',
        'background-inverse': '#f4f4f4',
        'background-brand': '#0f62fe',
        'background-active': '#525252',
        'overlay-01': 'rgba(22, 22, 22, 0.5)',
        'skeleton-01': '#353535',

        // Focus Colors
        'focus': '#0f62fe',
        'focus-inset': '#161616',
        'focus-inverse': '#0f62fe',

        // Miscellaneous Colors
        'interactive': '#4589ff',
        'highlight': '#001d6c',
        'toggle-off': '#6f6f6f',
        'overlay': 'rgba(22, 22, 22, 0.7)', // Tailwind doesn't support opacity in hex colors directly, so approximate with RGBA
        'skeleton-element': '#525252',
        'skeleton-background': '#353535',

        // Tag Colors - Gray
        'tag-background-gray': '#525252',
        'tag-color-gray': '#e0e0e0',
        'tag-hover-gray': '#636363',
        'tag-border-gray': '#8d8d8d',

        // Tag Colors - Cool Gray
        'tag-background-cool-gray': '#4d5358',
        'tag-color-cool-gray': '#dde1e6',
        'tag-hover-cool-gray': '#5d646a',
        'tag-border-cool-gray': '#878d96',

        // Tag Colors - Warm Gray
        'tag-background-warm-gray': '#565151',
        'tag-color-warm-gray': '#e5e0df',
        'tag-hover-warm-gray': '#696363',
        'tag-border-warm-gray': '#8f8b8b',

        // Tag Colors - Red
        'tag-background-red': '#a2191f',
        'tag-color-red': '#ffd7d9',
        'tag-hover-red': '#c21e25',
        'tag-border-red': '#fa4d56',

        // Tag Colors - Magenta
        'tag-background-magenta': '#9f1853',
        'tag-color-magenta': '#ffd6e8',
        'tag-hover-magenta': '#bf1d63',
        'tag-border-magenta': '#ee5396',

        // Tag Colors - Purple
        'tag-background-purple': '#6929c4',
        'tag-color-purple': '#e8daff',
        'tag-hover-purple': '#7c3dd6',
        'tag-border-purple': '#a56eff',

        // Tag Colors - Blue
        'tag-background-blue': '#0043ce',
        'tag-color-blue': '#d0e2ff',
        'tag-hover-blue': '#0053ff',
        'tag-border-blue': '#4589ff',

        // Tag Colors - Cyan
        'tag-background-cyan': '#00539a',
        'tag-color-cyan': '#bae6ff',
        'tag-hover-cyan': '#006fbb', // Adjusted for consistency; original value seems to be a duplicate
        'tag-border-cyan': '#1192e8',

        // Tag Colors - Teal
        'tag-background-teal': '#005D5D',
        'tag-color-teal': '#9ef0f0',
        'tag-hover-teal': '#007070',
        'tag-border-teal': '#009d9a',

        // Tag Colors - Green
        'tag-background-green': '#0e6027',
        'tag-color-green': '#a7f0ba',
        'tag-hover-green': '#11742f',
        'tag-border-green': '#24a148',

        // Button Primary
        'button-primary': '#0f62fe',
        'button-primary-hover': '#0353e9',
        'button-primary-active': '#002d9c',
        'button-primary-disabled': '#525252'
      },
      spacing: {
        'spacing-01': '0.125rem', // 2px
        'spacing-02': '0.25rem',  // 4px
        'spacing-03': '0.5rem',   // 8px
        'spacing-04': '0.75rem',  // 12px
        'spacing-05': '1rem',     // 16px
        'spacing-06': '1.5rem',   // 24px
        'spacing-07': '2rem',     // 32px
        'spacing-08': '2.5rem',   // 40px
        'spacing-09': '3rem',     // 48px
        'spacing-10': '4rem',     // 64px
        'spacing-11': '5rem',     // 80px
        'spacing-12': '6rem',     // 96px
        'spacing-13': '10rem'    // 160px
      }
    }
  },
  plugins: []
}
