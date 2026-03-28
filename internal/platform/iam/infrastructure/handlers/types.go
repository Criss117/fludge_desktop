package handlers

import "desktop/internal/appstate"

type OnStateChange func(e appstate.StateChangeEvent)
