$(document).ready(() => {
  $("#toggle-mobile-menu").on("click", () => {
    const menu = $("#mobile-menu");
    const currentMaxHeight = menu.css("max-height");

    if (currentMaxHeight === "0px" || currentMaxHeight === "none") {
      menu.css("max-height", "154px");
    } else {
      menu.css("max-height", "0px");
    }
  });
});