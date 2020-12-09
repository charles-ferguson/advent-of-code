#!/usr/bin/env ruby

class Program
  attr_reader :instructions, :instruction_visits, :positions
  attr_accessor :accumulator
  def initialize(data)
    @instructions = data.lines.map(&:chomp)
    @instruction_visits = Array.new(instructions.size, false)
    @accumulator = 0
    @positions = [0]
  end

  def position
    positions.last
  end

  def run
    while position < instructions.size
      return -1 if instruction_visits[position]

      instruction_visits[position] = true
      instruction = instructions[position]
      Instructions.for(instruction).execute_in(self)
    end

    0
  end
end

module Instructions
  INSTRUCTION_REGEX = /^(?<instruction>\S+) (?<operand>[+-]\d+)$/
  def self.for(line)
    match = line.match(INSTRUCTION_REGEX)
    raise "Invalid instruction format: #{line}" unless match

    operand = match[:operand].to_i
    case match[:instruction]
    when 'acc' then AccInstruction.new(operand)
    when 'nop' then NopInstruction.new(operand)
    when 'jmp' then JmpInstruction.new(operand)
    else
      raise "Instruction Type not implemented: #{line}"
    end
  end

  class NopInstruction
    attr_reader :operand
    def initialize(operand)
      @operand = operand
    end

    def execute_in(context)
      context.positions << context.position + 1
    end
  end

  class AccInstruction < NopInstruction
    def execute_in(context)
      context.accumulator += operand
      context.positions << context.position + 1
    end
  end

  class JmpInstruction < NopInstruction
    def execute_in(context)
      context.positions << context.position + operand
    end
  end
end

class LoopFixer
  ALLOWED_SWAPS = {
    'nop' => 'jmp',
    'jmp' => 'nop'
  }.freeze

  attr_reader :data
  def initialize(data)
    @data = data.lines
    @potential_swaps = @data.each_with_index.inject([]) do |collection, (instruction, index)|
      ALLOWED_SWAPS.keys.each do |swap|
        if instruction.start_with?(swap)
          collection << { position: index, swap: instruction.gsub(swap, ALLOWED_SWAPS.fetch(swap)) }
        end
      end

      collection
    end
  end

  def attempt_to_fix
    @potential_swaps.each do |swap|
      altered_program = data.dup
      altered_program[swap.fetch(:position)] = swap.fetch(:swap)

      program = Program.new(altered_program.join(''))
      return program if program.run == 0
    end

    raise 'No Simple instruction swap was able to fix the program'
  end
end



if $PROGRAM_NAME =~ /solution.rb$/
  data = File.read(File.join(__dir__, 'data'))
  program = Program.new(data)

  program.run
  puts "Part 1: #{program.accumulator}"
  loop_fixer = LoopFixer.new(data)
  program = loop_fixer.attempt_to_fix
  puts "Part 2: #{program.accumulator}"
end
