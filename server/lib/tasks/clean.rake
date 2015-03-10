task :clean => :environment do
	CarrierWave.clean_cached_files!
end
