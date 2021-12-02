INPUT_FILE= File.join(__dir__, "data")

class Depths
  def self.from_file(file_path = INPUT_FILE)
    depths = []
    File.open(file_path).each do |line|
      depths << line.to_i
    end
    new(depths)
  end

  attr_reader :depths
  def initialize(depths)
    @depths = depths.dup
  end

  def count_increases
    count = 0
    depths.each.with_index { |depth, index| count += 1 if index != 0 && depths[index - 1] < depth }
    count
  end

  def windows(size = 3)
    windows = Array(0..(depths.count - 2)).map { |i| depths.slice(i, 3).reduce(:+) }
    self.class.new(windows)
  end
end

class Part1
  def self.answer
    depths = Depths.from_file
    depths.count_increases
  end
end

class Part2
  def self.answer
    depths = Depths.from_file
    depths.windows.count_increases
  end
end

puts Part1.answer
puts Part2.answer
