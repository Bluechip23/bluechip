#!/bin/bash

# debug
# set -x

# Function to check if screen is installed
check_screen_installed() {
    if ! command -v screen &> /dev/null; then
        echo "Error: screen is not installed. Please install screen and try again."
        exit 1
    fi
}

# Function to prompt user to stop existing screen sessions
prompt_to_stop_screens() {
    local screen_name=$1
    local existing_screens=()
    screen -ls "$screen_name" | grep -q "$screen_name"
    if [ $? -eq 0 ]; then
        existing_screens+=("$screen_name")
    fi

    if [[ ${#existing_screens[@]} -gt 0 ]]; then
        echo "The following screens already exist:"
        for screen_name in "${existing_screens[@]}"; do
            echo "- $screen_name"
        done

        read -p "Do you want to stop these screens? (y/n): " choice
        case $choice in
            [Yy]*)
                for screen_name in "${existing_screens[@]}"; do
                    screen -XS "${screen_name}" stuff "^C"
                    screen -XS "${screen_name}" quit
                done
                ;;
            [Nn]*)
                echo "Exiting..."
                exit 0
                ;;
            *)
                echo "Invalid choice. Exiting..."
                exit 1
                ;;
        esac
    fi
}

# Function to check if port is open
check_port_open() {
  local ip=$1
  local port=$2
  nc -z "$ip" "$port" >/dev/null 2>&1
}

# Wait for ports to become available
wait_for_ports_running() {
    local ip=$1
    local port=$2

    local timeout=30
    local interval=5
    local counter=0

    echo "Waiting for RPC ports to become available..."

    until check_port_open "$ip" "$port" || [ "$counter" -eq "$timeout" ]; do
        counter=$((counter + interval))
        sleep "$interval"
    done

    if [ "$counter" -eq "$timeout" ]; then
        echo "Timeout: Failed to connect to RPC port $port."
        exit 1
    fi

    echo "RPC port $port is now open."

    echo "Waiting 5 seconds for blockchain state to become available..."
    sleep 5
}

# Function to fetch node ID given an IP
fetch_node_id() {
    local node_ip="$1"
    local port=$2
    curl "http://${node_ip}:${port}/p2pid" || echo ""
}

# Function to add a peer to the bootstrap peers string
add_bootstrap_peer() {
    local node_ip="$1"
    local node_id="$2"
    if [ -n "$node_id" ]; then
        # Append the peer info to the bootstrap_peers string
        # Each peer is added as a quoted string
        bootstrap_peers="${bootstrap_peers}\"/ip4/${node_ip}/tcp/5051/ipfs/${node_id}\","
        bootstrap_peers_eddsa="${bootstrap_peers_eddsa}\"/ip4/${node_ip}/tcp/5052/ipfs/${node_id}\","
    fi
}