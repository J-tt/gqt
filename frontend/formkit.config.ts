// formkit.config.js
import { generateClasses } from '@formkit/themes'
import { genesisIcons } from '@formkit/icons'
import tailwindTheme from './formkit.theme' // change to your theme's path

export default {
  icons: {
    ...genesisIcons,
  },
  config: {
    classes: generateClasses(tailwindTheme),
  },
}
