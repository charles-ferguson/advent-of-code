#!/usr/bin/env ruby

INPUT_FILE = File.join(__dir__, 'data')
SAMPLE_FILE = File.join(__dir__, 'sample_data')

Instruction = Struct.new(:size)
Position = Struct.new(:depth, :horizontal_position, :aim)
class Position
  def output
    depth * horizontal_position
  end
end

module Part1
  class Forward < Instruction
    def move(original_position)
      position = Position.new(original_position.depth, original_position.horizontal_position + size)
    end
  end

  class Up < Instruction
    def move(original_position)
      position = Position.new(original_position.depth - size, original_position.horizontal_position)
    end
  end

  class Down < Instruction
    def move(original_position)
      position = Position.new(original_position.depth + size, original_position.horizontal_position)
    end
  end

  def self.run
    instructions = File.read(INPUT_FILE).lines.map { |l| Instructions.from_line(l, Part1) }
    submarine = Submarine.new
    instructions.each { |i| submarine.perform(i) }
    puts submarine.position.output
  end
end

module Part2
  class Forward < Instruction
    def move(original_position)
      position = original_position.dup
      position.depth += original_position.aim * size
      position.horizontal_position += size
      position
    end
  end

  class Up < Instruction
    def move(original_position)
      position = original_position.dup
      position.aim -= size
      position
    end
  end

  class Down < Instruction
    def move(original_position)
      position = original_position.dup
      position.aim += size
      position
    end
  end

  def self.run
    instructions = File.read(INPUT_FILE).lines.map { |l| Instructions.from_line(l, Part2) }
    submarine = Submarine.new
    instructions.each { |i| submarine.perform(i) }
    puts submarine.position.output
  end
end

class Instructions
  PART1_INSTRUCTION_MAP = {
    'forward' => Part1::Forward,
    'down'    => Part1::Down,
    'up'      => Part1::Up,
  }.freeze

  PART2_INSTRUCTION_MAP = {
    'forward' => Part2::Forward,
    'down'    => Part2::Down,
    'up'      => Part2::Up,
  }.freeze

  def self.from_line(line, part = Part1)
    data = line.match(/^(?<direction>\w+) (?<size>\d+)\s*$/)
    raise "No known instruction for #{line}" if !data

    if part == Part1
      PART1_INSTRUCTION_MAP.fetch(data[:direction]).new(data[:size].to_i)
    elsif part == Part2
      PART2_INSTRUCTION_MAP.fetch(data[:direction]).new(data[:size].to_i)
    end
  end
end

class Submarine
  attr_reader :position
  def initialize(position = Position.new(0, 0, 0))
    @position = position
  end

  def perform(instruction)
    @position = instruction.move(position)
  end
end

Part1.run
Part2.run
