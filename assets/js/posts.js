$("#new-post").on("submit", CreatePost);

$(document).on("click", ".like-post", LikePost);
$(document).on("click", ".dislike-post", DislikePost);

$("#update-post").on("click", UpdatePost);
$(".delete-post").on("click", DeletePost);

function CreatePost(event) {
  event.preventDefault();

  $.ajax({
    url: "/posts",
    method: "POST",
    data: {
      title: $("#title").val(),
      content: $("#content").val(),
    },
  })
    .done(() => {
      window.location = "/home";
    })
    .fail(() => {
      Swal.fire("Error", "Error in post creation", "error");
    });
}

function LikePost(event) {
  event.preventDefault();

  const clickedElement = $(event.target);
  const postId = clickedElement.closest("div").data("post-id");
  clickedElement.prop("disabled", true);

  $.ajax({
    url: `/posts/${postId}/like`,
    method: "POST",
  })
    .done(() => {
      const likesElement = clickedElement.next("span");
      const likes = parseInt(likesElement.text());

      likesElement.text(likes + 1 + " likes");

      clickedElement.addClass("dislike-post");
      clickedElement.css("color", "red");
      clickedElement.removeClass("like-post");
    })
    .fail(() => {
      Swal.fire("Error", "Error while liking post", "error");
    })
    .always(() => {
      clickedElement.prop("disabled", false);
    });
}

function DislikePost(event) {
  event.preventDefault();

  const clickedElement = $(event.target);
  const postId = clickedElement.closest("div").data("post-id");

  clickedElement.prop("disabled", true);

  $.ajax({
    url: `/posts/${postId}/dislike`,
    method: "POST",
  })
    .done(() => {
      const likesElement = clickedElement.next("span");
      const likes = parseInt(likesElement.text());

      likesElement.text(likes - 1 + " likes");

      clickedElement.removeClass("dislike-post");
      clickedElement.css("color", "");
      clickedElement.addClass("like-post");
    })
    .fail(() => {
      Swal.fire("Error", "Error while disliking post", "error");
    })
    .always(() => {
      clickedElement.prop("disabled", false);
    });
}

function UpdatePost(event) {
  event.preventDefault();

  $(this).prop("disabled", true);

  const postId = $(this).data("post-id");

  $.ajax({
    url: `/posts/${postId}`,
    method: "PUT",
    data: {
      title: $("#title").val(),
      content: $("#content").val(),
    },
  })
    .done(() => {
      Swal.fire("Success", "Post updated successfully", "success").then(() => (window.location = "/home"));
    })
    .fail(() => {
      Swal.fire("Error", "Error in post update", "error");
    })
    .always(() => {
      $("#updatePost").prop("disabled", false);
    });
}

function DeletePost(event) {
  event.preventDefault();

  Swal.fire({
    title: "Warning!",
    text: "Are you sure you want to delete this post?",
    showCancelButton: true,
    cancelButtonText: "Cancel",
    icon: "warning",
  }).then((confirmation) => {
    if (!confirmation.value) return;

    const clickedElement = $(event.target).closest("div");
    const postId = clickedElement.data("post-id");

    clickedElement.prop("disabled", true);

    $.ajax({
      url: `/posts/${postId}`,
      method: "DELETE",
    })
      .done(() => {
        $("#user-post").fadeOut("slow", () => $(this).remove());
      })
      .fail(() => {
        Swal.fire("Error", "Error on post deletion", "error");
      })
      .always(() => {
        $("#updatePost").prop("disabled", false);
      });
  });
}
