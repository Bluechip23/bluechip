// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import Smartdev0328BluechipBluechipMint from './smartdev0328/bluechip/bluechip.mint'
import Smartdev0328BluechipSmartdev0328BluechipPageinflation from './smartdev0328/bluechip/smartdev0328.bluechip.pageinflation'


export default { 
  Smartdev0328BluechipBluechipMint: load(Smartdev0328BluechipBluechipMint, 'bluechip.mint'),
  Smartdev0328BluechipSmartdev0328BluechipPageinflation: load(Smartdev0328BluechipSmartdev0328BluechipPageinflation, 'smartdev0328.bluechip.pageinflation'),
  
}


function load(mod, fullns) {
    return function init(store) {        
        if (store.hasModule([fullns])) {
            throw new Error('Duplicate module name detected: '+ fullns)
        }else{
            store.registerModule([fullns], mod)
            store.subscribe((mutation) => {
                if (mutation.type == 'common/env/INITIALIZE_WS_COMPLETE') {
                    store.dispatch(fullns+ '/init', null, {
                        root: true
                    })
                }
            })
        }
    }
}
