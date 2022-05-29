const webAnalytics = "webAnalytics";
let userToken;

window.addEventListener('load', sendEvent, false)


function sendEvent() {
    let cookie = getCookie(webAnalytics)
    let visit = {
        id: cookie,
        url: window.location.href
    }
    fetch("http://localhost:8080/events/visit", {
        method: 'POST',
        body: JSON.stringify(visit),
        headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json'
        }
    }).then(function (response) {
        return response.json();
    }).then(token => {
        userToken = token.id;
        if(cookie === '') {
            setCookie(webAnalytics, token.id, 30);
        }
        let event = {
            userAgent: navigator.userAgent,
            screenWight: window.screen.width,
            screenHeight: window.screen.height,
            browserLanguage: navigator.language,
            browserOnline: navigator.onLine,
            visitId: userToken,
        }
        fetch("http://localhost:8080/events",
            {
                method: 'POST',
                body: JSON.stringify(event),
                headers: {
                    'Accept': 'application/json',
                    'Content-Type': 'application/json'
                }
            }).then(res => {
            console.log("Request complete! response:", res);
            }).catch(e => {
            console.error("Request Error: ", e);
            });
    })

}

function setCookie(name, value, exp_days) {
    let d = new Date();
    d.setTime(d.getTime() + (exp_days*24*60*60*1000));
    let expires = "expires=" + d.toUTCString;
    document.cookie = name + "=" + value + ";" + expires + ";path=/";
}

function getCookie(name) {
    let cname = name + "=";
    let decodedCookie = decodeURIComponent(document.cookie)
    let ca = decodedCookie.split(';');
    for(let i = 0; i < ca.length; i++) {
        let c = ca[i];
        while (c.charAt(0) === ' ') {
            c = c.substring(1);
        }
        if(c.indexOf(cname) === 0) {
            return c.substring(cname.length, c.length)
        }
    }
    return "";
}
