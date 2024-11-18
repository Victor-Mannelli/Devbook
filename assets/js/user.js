$("#unfollow").on("click", Unfollow);
$("#follow").on("click", Follow);

function Unfollow() {
  const userId = $(this).data("user-id");
  $(this).prop("disabled", true);

  $.ajax({
    url: `/users/${userId}/unfollow`,
    method: "Post",
  })
    .done(() => (window.location = `/users/${userId}`))
    .fail(() => {
      Swal.fire("Error", "Error in unfollowing user", "error");
      $("#unfollow").prop("disabled", false);
    });
}

function Follow() {
  const userId = $(this).data("user-id");
  $(this).prop("disabled", true);

  $.ajax({
    url: `/users/${userId}/follow`,
    method: "Post",
  })
    .done(() => (window.location = `/users/${userId}`))
    .fail(() => {
      Swal.fire("Error", "Error in following user", "error");
      $("#follow").prop("disabled", false);
    });
}
