#!/usr/bin/env ruby

DATA = File.read(File.join(__dir__, "data"))

class Map
  EMPTY = "."
  TREE = "#"

  attr_reader :map
  def initialize(map)
    @map = read_map(map)
  end

  def tree?(position)
    wrapped_position = wrap_position(position)
    map.fetch(wrapped_position.y).fetch(wrapped_position.x) == TREE
  end

  def off_map?(position)
    position.y >= map.size
  end

  def print_with(positions)
    map_to_mark = clone_map
    positions.each do |position|
      wrapped_position = wrap_position(position)
      map_to_mark[wrapped_position.y][wrapped_position.x] = tree?(wrapped_position) ? 'X' : 'O'
    end

    map_to_mark.each { |row| puts row.inspect }
  end

  def clone_map
    clone = []
    map.each { |row| clone << row.dup }
    clone
  end

  private

  def wrap_position(position)
    Point.new(position.y, position.x % map.first.size)
  end

  def read_map(string_map)
    string_map.lines.inject([]) do |array_map, line|
      array_map << line.chomp.split('')
    end
  end
end

Point = Struct.new(:y, :x)
class Point
  def move(down, right)
    Point.new(y + down, x + right)
  end
end

def trees_on_slope(slope, world)
  locations = [Point.new(*[0,0])]
  loop do
    next_location = locations.last.move(*slope)
    break if world.off_map?(next_location)

    locations << next_location
  end
  locations.select { |loc| world.tree?(loc) }.count
end

world = Map.new(DATA)
puts "Part 1: trees hit on initial slope #{trees_on_slope([1, 3], world)}"

all_slopes = [[1, 1], [1, 3], [1, 5], [1, 7], [2, 1]]
counts = all_slopes.map { |slope| trees_on_slope(slope, world) }
puts "Part 2: product of trees hit on all slopes #{counts.reduce(:*)}"
