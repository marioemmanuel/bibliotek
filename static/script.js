// PROJECT: bibliotek makefile
// AUTHOR: MARIO EMMANUEL
// DATE: 2023/DEC/19
// github.com/marioemmanuel/bibliotek

document.addEventListener("DOMContentLoaded", function() {
    fetch('/tree')
        .then(response => response.text())
        .then(html => {
            var navPane = document.getElementById('navPane');
            navPane.innerHTML = html;

            // Set up caret click handlers after the tree is loaded
            var folders = navPane.getElementsByClassName("folder-item");
            for (var i = 0; i < folders.length; i++) {
                folders[i].addEventListener("click", function() {

			        // Retrieving child elements
			        var caret = this.querySelector('.caret');
			        var folderCloseIcon = this.querySelector('.folderclose');
			        var folderOpenIcon = this.querySelector('.folderopen');
	              	var nestedList = this.nextElementSibling;

					// Toggle caret 
					caret.classList.toggle("caret-down");

					// Search for next nested class element (the ul)
					if (nestedList !== null) {
                    	if (nestedList.style.display === "block") {
                    	    folderOpenIcon.style.display = "none";
                    	    folderCloseIcon.style.display = "inline-block";
                    	    nestedList.style.display = "none";
                    	} else {
                    	    folderOpenIcon.style.display = "inline-block";
                    	    folderCloseIcon.style.display = "none";
                    	    nestedList.style.display = "block";
                    	}
					}
                });
            }

            // Set up page click handlers
            var pageLinks = navPane.querySelectorAll("a[content-path]");
            pageLinks.forEach(link => {
                link.addEventListener("click", function(event) {
                    var mdPath = this.getAttribute('content-path');
                    fetch('/content' + mdPath) // Adjust the endpoint as needed
                        .then(response => response.text())
                        .then(htmlContent => {
                            document.getElementById('contentPane').innerHTML = htmlContent;
                        })
                        .catch(err => console.error('Failed to load content:', err));
                });
            });

        })
        .catch(err => console.error('Failed to load tree:', err));
});
