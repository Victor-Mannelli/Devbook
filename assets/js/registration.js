$("#signup-form").on("submit", createUser);

function createUser(event) {
  event.preventDefault();

  if ($("#password").val() != $("#confirm_password").val()) {
    alert("different passwords")
  }

  $.ajax({
    url: "/users",
    method: "POST",
    data: {
      name: $('#name').val(),
      email: $('#email').val(),
      username: $('#username').val(),
      password: $('#password').val(),
    } 
  }).done(() => {
    alert("user created")
  }).fail((err) => {
    console.log(err)
    alert("error at user creation")
  })
}
