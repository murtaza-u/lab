const image = document.querySelector("img");
const btn = document.querySelector("button");
const input = document.querySelector("input");

function getRoundedCanvas(sourceCanvas) {
  var canvas = document.createElement("canvas");
  var context = canvas.getContext("2d");
  var width = sourceCanvas.width;
  var height = sourceCanvas.height;

  canvas.width = width;
  canvas.height = height;
  context.imageSmoothingEnabled = true;
  context.drawImage(sourceCanvas, 0, 0, width, height);
  context.globalCompositeOperation = "destination-in";
  context.beginPath();
  context.arc(
    width / 2,
    height / 2,
    Math.min(width, height) / 2,
    0,
    2 * Math.PI,
    true,
  );
  context.fill();
  return canvas;
}

let croppable = false;
let cropper;

btn.addEventListener("click", () => {
  if (!croppable) return;

  const croppedCanvas = cropper.getCroppedCanvas();

  const roundedCanvas = getRoundedCanvas(croppedCanvas);
  roundedImage = document.createElement("img");
  roundedImage.src = roundedCanvas.toDataURL();
  roundedImage.style = "width: 360px;";

  createFile(roundedImage.src);
});

async function createFile(url) {
  const response = await fetch(url);
  const data = await response.blob();
  const file = new File([data], "avatar.png", { type: data.type });
  const dt = new DataTransfer();
  dt.items.add(file);
  input.files = dt.files;
  const event = new Event("change", {
    bubbles: !0,
  });
  input.dispatchEvent(event);
}

input.addEventListener("change", (e) => {
  if (cropper) {
    cropper.destroy();
    cropper = null;
  }

  const files = e.target.files;
  if (!files || files.length === 0) {
    return;
  }

  const done = (url) => {
    input.value = null;
    image.src = url;
    cropper = new Cropper(image, {
      aspectRatio: 1,
      dragMode: "move",
      cropBoxMovable: true,
      cropBoxResizable: false,
      toggleDragModeOnDblclick: false,
      responsive: true,
      zoomable: false,
      center: true,
      data: {
        width: 360,
        height: 360,
      },
      ready: () => {
        croppable = true;
      },
    });
  };

  const file = files[0];

  if (URL) {
    done(URL.createObjectURL(file));
    return;
  }

  if (FileReader) {
    const reader = new FileReader();
    reader.onload = () => done(reader.result);
    reader.readAsDataURL(file);
    return;
  }
});
