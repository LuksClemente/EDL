import Math;
public class Main{

	public interface Geometria{
		public float area();
		public float perim();
	}

	public static void main(String[] args){
		Retangulo a = new Retangulo(3,4);
		Circulo b = new Circulo(5);
		System.out.println(a.area());
		System.out.println(a.perim());
		System.out.println(b.area());
		System.out.println(b.perim());
	}
}

public class Retangulo implements Geometria{
	private float altura;
	private float largura;

	public float area(float largura, float altura){
		return this.largura*this.altura;
	}

	public float perim(float largura, float altura){
		return 2*this.largura + 2*this.altura;
	}
}

public class Circulo implements Geometria{
	private float raio;

	public float area(float raio){
	return Math.PI * this.raio * this.raio;
	}

	public float perim(float raio){
		return 2*Math.PI*this.raio;
	}
}
