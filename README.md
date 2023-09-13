# rai-snowflake-demo-app

This repo provides a server that folks can use to run the the Snowflake-RelationalAI demo locally.

## Getting Started

### Downloading

*Note: this application delivers Act 2 (Graph Visualization) and Act 3 (Optimization). If you only want Act 2, first navigate to this branch and then continue with the instructions below: [https://github.com/sswatson/rai-snowflake-demo-app/tree/act-2-only](https://github.com/sswatson/rai-snowflake-demo-app/tree/act-2-only).*

Start by clicking the green "Code" button above and select "Download ZIP" to download this repo. Unzip the directory on your local machine.

### Running the demo - Windows

1. Open a command prompt.
2. Navigate to the unzipped folder location.
3. In the command prompt type: `server.exe`.
4. Hit `Enter`.
5. Open the following url in your browser:
   - localhost:8080

### Running the demo - Mac OSX

1. Open a terminal window.
2. Navigate to the unzipped folder location (`cd ~/Downloads/rai-snowflake-demo-app-main`).
3. In the terminal type: `chmod +x server`. Press Enter.
4. In the terminal type: `xattr -dr com.apple.quarantine server`. Press Enter.
5. Hit `Enter`.
6. In the terminal type: `./server`.
7. Hit `Enter`.
8. Open the following url in your browser:
   - localhost:8080
