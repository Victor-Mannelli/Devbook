$("#signup-form").on("submit", createUser);

function createUser(event) {
  event.preventDefault();

  if ($("#password").val() != $("#confirm_password").val()) {
    Swal.fire("Error", "Passwords don't match", "error");
    return
  }

  $.ajax({
    url: "/users",
    method: "POST",
    data: {
      name: $("#name").val(),
      email: $("#email").val(),
      username: $("#username").val(),
      password: $("#password").val(),
    },
  })
    .done(() => {
      Swal.fire("Success!", "User created successfully", "success");
    })
    .fail(() => {
      Swal.fire("Error", "Error at user creation", "error");
    });
}
