let hideHeader = true, keepHeaderShowingStart = true;

function headerHoverToggle() {
    if (keepHeaderShowingStart) {
        document.getElementById("website-header").classList.add("header-autohide");
        keepHeaderShowingStart = false;
    }
}

function headerToggleHide() {
    keepHeaderShowingStart = false;
    if (hideHeader) {
        document.getElementById("toggle-header-button").innerHTML = "<i class=\"material-icons\">lens</i>";
        document.getElementById("website-header").classList.remove("header-autohide");
        hideHeader = false;
    } else {
        document.getElementById("toggle-header-button").innerHTML = "<i class=\"material-icons\">panorama_fish_eye</i>";
        document.getElementById("website-header").classList.add("header-autohide");
        hideHeader = true;
        // messy headerMouseLeave();
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