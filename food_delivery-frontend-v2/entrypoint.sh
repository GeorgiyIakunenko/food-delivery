#!/bin/sh
set -e

# Install Node packages
npm install

# Start the app
npm run build && npm run preview
```
