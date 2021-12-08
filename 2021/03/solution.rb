module InputData
  INPUT_FILE = File.join(__dir__, "data")
  SAMPLE_FILE = File.join(__dir__, "sample_data")

  def self.get_readings(file)
    readings = Readings.new
    File.open(file).each { |line| readings.add(line) }
    readings
  end
end

class Readings
  attr_reader :list
  def initialize
    @list = []
    @counts = []
  end

  def add(line)
    line = line.chomp
    list << line
    line.each_char.with_index do |char, index|
      @counts[index] = Hash.new { |hash, key| hash[key] = 0 } if @counts[index].nil?
      @counts[index][char] += 1
    end
  end

  def gamma_rating
    binary_rating = @counts.each.with_object("") do |count, rating|
      char = count.max_by { |_, v| v }
      rating << char.first
    end
    binary_rating.to_i(2)
  end

  def epsilon_rating
    binary_rating = @counts.each.with_object("") do |count, rating|
      char = count.min_by { |_, v| v }
      rating << char.first
    end
    binary_rating.to_i(2)
  end

  def oxygen_rating
    binary_rating = @counts.each_with_index.inject(list.dup) do |potential_list, (count, index)|
      break potential_list if potential_list.size == 1

      split = potential_list.group_by { |reading| reading[index] }
      potential_list = case split.fetch("0", []).size <=> split.fetch("1", []).size
        when -1
          split.fetch("1")
        when 1
          split.fetch("0")
        when 0
          split.fetch("1")
        end
    end

    binary_rating.first.to_i(2)
  end

  def co2_rating
    binary_rating = @counts.each_with_index.inject(list.dup) do |potential_list, (count, index)|
      break potential_list if potential_list.size == 1

      split = potential_list.group_by { |reading| reading[index] }
      potential_list = case split.fetch("0", []).size <=> split.fetch("1", []).size
        when -1
          split.fetch("0")
        when 1
          split.fetch("1")
        when 0
          split.fetch("0")
        end
    end

    binary_rating.first.to_i(2)
  end

  def life_support_rating
    oxygen_rating * co2_rating
  end

  def power_rating
    epsilon_rating * gamma_rating
  end
end

module Part1
  def self.run
    readings = InputData.get_readings(InputData::INPUT_FILE)
    readings.power_rating
  end
end

module Part2
  def self.run
    readings = InputData.get_readings(InputData::INPUT_FILE)
    readings.life_support_rating
  end
end

puts Part1.run
puts  Part2.run
