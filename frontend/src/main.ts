import { createApp } from 'vue'
import App from './App.vue'
import { plugin, defaultConfig } from '@formkit/vue'
import formkitConfig from '../formkit.config'
import './style.css';

createApp(App).use(plugin, defaultConfig(formkitConfig)).mount('#app')


