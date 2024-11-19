$("#unfollow").on("click", Unfollow);
$("#follow").on("click", Follow);
$("#edit-user").on("submit", EditUser);
$("#update-password").on("submit", UpdatePassword);

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

function EditUser(event) {
  event.preventDefault();

  $.ajax({
    url: "/edit-user",
    method: "PUT",
    data: {
      name: $("#name").val(),
      email: $("#email").val(),
      username: $("#username").val(),
    },
  })
    .done(() => {
      Swal.fire("Success", "User updated successfully", "success").then(() => (window.location = "/profile"));
    })
    .fail(() => {
      Swal.fire("Error", "Error in user update", "error");
    });
}

function UpdatePassword(event) {
  event.preventDefault();

  if ($("#new-password").val() != $("#confirm-password").val()) {
    Swal.fire("Warning!", "Password don't match", "warning");
    return
  }
  
  $.ajax({
    url: "/update-password",
    method: "PUT",
    data: {
      password: $("#current-password").val(),
      newPassword: $("#new-password").val(),
    },
  })
    .done(() => {
      Swal.fire("Success", "Password updated successfully", "success").then(() => (window.location = "/profile"));
    })
    .fail((e) => {
      console.log(e)
      Swal.fire("Error", "Error in password update", "error");
    });
}
