function generateUUID() {
    var uuid = 'xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx'.replace(/[xy]/g, function(c) {
                var r = Math.random() * 16 | 0, v = c == 'x' ? r : (r & 0x3 | 0x8);
                return v.toString(16);
            
    });
        document.getElementById("system-uuid").value = uuid;
    
}

function saveData() {
        var systemName = document.getElementById("system-name").value;
        var systemId = document.getElementById("system-id").value;
        var systemUUID = document.getElementById("system-uuid").value;
        var image = canvas.toDataURL("image/png");
    // sending data to the middleware
    var data = {
                systemName: systemName,
                systemId: systemId,
                systemUUID: systemUUID,
                image: image
            
    }

        // using fetch API to send data to the middleware
    fetch('/middleware', {
                method: 'POST',
                body: JSON.stringify(data),
        headers: {
                        'Content-Type': 'application/json'
                    
        }
    }).then(res => res.json())
        .then(response => {
                    console.log('Success:', JSON.stringify(response));
                
        })
        .catch(error => {
                    console.error('Error:', error);
                
        });
}
