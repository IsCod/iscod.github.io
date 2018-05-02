$(document).ready(function(){
	$.get("/public/header.html", function(data,status){
		$("body").prepend(data);
	})

	$.get("/public/footer.html", function(data,status){
		$("body").append(data);
	})
});