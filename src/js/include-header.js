let hideHeader = true;

function headerToggleHide() {
    if (hideHeader) {
        document.getElementById("toggle-header-button").innerHTML = "<i class=\"material-icons\">lens</i>";
        headerMouseEnter();
        hideHeader = false;
    } else {
        document.getElementById("toggle-header-button").innerHTML = "<i class=\"material-icons\">panorama_fish_eye</i>";
        hideHeader = true;
        // messy headerMouseLeave();
    }
}

function headerMouseEnter() {
    if (hideHeader) {
        document.getElementById("site-header").classList.remove("header-hidden");
        document.getElementById("site-header").classList.add('header-visible');
    }
}

function headerMouseLeave() {
    if (hideHeader) {
        document.getElementById("site-header").classList.remove("header-visible");
        document.getElementById("site-header").classList.add('header-hidden');
    }
}

function includeHTML() {
    let z, i, elmnt, file, xhttp;
    /* Loop through a collection of all HTML elements: */
    z = document.getElementsByTagName("*");
    for (i = 0; i < z.length; i++) {
        elmnt = z[i];
        /*search for elements with a certain atrribute:*/
        file = elmnt.getAttribute("w3-include-html");
        if (file) {
            /* Make an HTTP request using the attribute value as the file name: */
            xhttp = new XMLHttpRequest();
            xhttp.onreadystatechange = function() {
                if (this.readyState == 4) {
                    if (this.status == 200) {elmnt.innerHTML = this.responseText;}
                    if (this.status == 404) {elmnt.innerHTML = "Page not found.";}
                    /* Remove the attribute, and call this function once more: */
                    elmnt.removeAttribute("w3-include-html");
                    includeHTML();
                }
            };
            xhttp.open("GET", file, true);
            xhttp.send();
            /* Exit the function: */
            return;
        }
    }
}
includeHTML();