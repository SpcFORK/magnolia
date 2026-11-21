package reference {
  import reference.Item;
  public class Main {
    public var initial:int = 17;
    public function Main() {}
    public function values():Array {
      var a:int = 1, b:String = "two", c:int = Math.max(2, 3);
      return [a, b, c];
    }
    public function regex():Boolean {
      var pattern:RegExp = /a{2}\/b/gi;
      return pattern.test("aa/b");
    }
    public function quotient():Number { return 8 / 2; }
    public function cast(value:*):String { return value as String; }
    public function isInt(value:*):Boolean { return value is int; }
    public function item():Item { return new Item(); }
    public function isItem(value:*):Boolean { return value is Item; }
  }
}
