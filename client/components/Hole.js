export default class Hole {
  constructor(props) {
    this.canvas = props.canvas;
    this.position = props.position;
    this.radius = 25;
    this.lifespan = 20000;
    this.getPositionAndRadius = this.getPositionAndRadius.bind(this);
    this.drawHole = this.drawHole.bind(this);
  }

  getPositionAndRadius() {
    return { position: this.position, radius: this.radius };
  }

  drawHole() {
    const ctx = this.canvas.getContext('2d');
    ctx.beginPath();
    ctx.arc(this.position.x, this.position.y, this.radius, 0, Math.PI * 2);
    ctx.fillStyle = 'white';
    ctx.fill();
    ctx.closePath();
  }
}
