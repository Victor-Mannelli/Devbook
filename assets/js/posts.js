$("#new-post").on("submit", CreatePost);

$(document).on("click", ".like-post", LikePost);
$(document).on("click", ".dislike-post", DislikePost);

$("#update-post").on("click", UpdatePost)

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

function LikePost(event) {
  event.preventDefault();

  const clickedElement = $(event.target);
  const postId = clickedElement.closest('div').data('post-id');

  clickedElement.prop('disabled', true)

  $.ajax({
    url: `/posts/${postId}/like`,
    method: "POST",
  }).done(() => {
    const likesElement = clickedElement.next('span');
    const likes = parseInt(likesElement.text());

    likesElement.text(likes + 1);

    clickedElement.addClass('dislike-post')
    clickedElement.css("color", "red");
    clickedElement.removeClass('like-post')

  }).fail((err) => {
    console.log(err)
    alert("error while liking post");
  }).always(() => {
    clickedElement.prop('disabled', false)
  })
}

function DislikePost(event) {
  event.preventDefault();

  const clickedElement = $(event.target);
  const postId = clickedElement.closest('div').data('post-id');

  clickedElement.prop('disabled', true)

  $.ajax({
    url: `/posts/${postId}/dislike`,
    method: "POST",
  }).done(() => {
    const likesElement = clickedElement.next('span');
    const likes = parseInt(likesElement.text());

    likesElement.text(likes - 1);

    clickedElement.removeClass('dislike-post')
    clickedElement.css("color", "");
    clickedElement.addClass('like-post')

  }).fail((err) => {
    console.log(err)
    alert("error while liking post");
  }).always(() => {
    clickedElement.prop('disabled', false)
  })
}

function UpdatePost(event) {
  event.preventDefault();

  $(this).prop("disabled", true);

  const postId = $(this).data('post-id');

  $.ajax({
    url: `/posts/${postId}`,
    method: "PUT",
    data: {
      title: $('#title').val(),
      content: $('#content').val(),
    } 
  }).done(() => {
    alert("success on post update");
  }).fail((err) => {
    alert("error on post update");
  }).always(() => {
    $("#updatePost").prop("disabled", false);
  })
};