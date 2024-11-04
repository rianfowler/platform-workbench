# platform-repo/Tiltfile

# Path to the services directory
SERVICES_DIR = 'services'

# List of service directories
services = [
    'other-api',
    'platform_api',
    'rians-cli'
]

# Iterate over each service and include its Tiltfile
for service in services:
    include(SERVICES_DIR + '/' + service + '/Tiltfile')