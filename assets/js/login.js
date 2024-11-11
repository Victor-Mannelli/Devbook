$("#login-form").on("submit", login);

function login(event) {
  event.preventDefault();

  $.ajax({
    url: "/login",
    method: "POST",
    data: {
      email: $('#email').val(),
      password: $('#password').val(),
    } 
  }).done(() => {
    window.location = "/home"
  }).fail((err) => {
    console.log(err)
    alert("user or password invalid")
  })
}
