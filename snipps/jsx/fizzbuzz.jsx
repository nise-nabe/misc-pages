class _Main {
  static function main(args : string[]) :void {
    for (var i = 0; i < 100; ++i) {
      if (i % 15 == 0) {
        log "fizzbuzz";
      } else if(i % 3 == 0) {
        log "fizz";
      } else if(i % 3 == 0) {
        log "buzz";
      } else {
        log i;
      }
    }
  }
}
