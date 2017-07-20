import injectTapEventPlugin from 'react-tap-event-plugin';

import {detectLocale} from './intl'
import main from './main'
import registerServiceWorker from './registerServiceWorker';

import './main.css';

injectTapEventPlugin()

const user = detectLocale()
main('root', user)
registerServiceWorker();
