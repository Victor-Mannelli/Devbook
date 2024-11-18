$("#login-form").on("submit", login);

function login(event) {
  event.preventDefault();

  $.ajax({
    url: "/login",
    method: "POST",
    data: {
      email: $("#email").val(),
      password: $("#password").val(),
    },
  })
    .done(() => {
      window.location = "/home";
    })
    .fail((err) => {
      Swal.fire("Error", "Email and Password combination are invalid", "error");
    });
}
