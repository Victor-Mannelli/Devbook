$("#new-post").on("submit", CreatePost);

function CreatePost(event) {
  event.preventDefault();

  $.ajax({
    url: "/posts",
    method: "POST",
    data: {
      title: $('#title').val(),
      content: $('#content').val(),
    } 
  }).done(() => {
    console.log("test")
    window.location = "/home";
  }).fail((err) => {
    console.log(err)
    alert("error on post creation");
  })
}